# Chạy ICM relayer cho bridge Sepolia ⇄ Fuji

Hợp đồng đã deploy sẵn trên cả hai mạng. Infra chỉ cần chạy **3 service** trên **1 server Linux**.

| Service | Làm gì |
|---|---|
| `avalanchego-fuji` | Node Fuji (chỉ P-Chain), để relayer lấy danh sách validator |
| `lc-updater` | Cập nhật ZK light client của Sepolia trên Fuji (tạo proof) |
| `icm-relayer` | Chuyển message hai chiều Sepolia ⇄ Fuji |

Node Sepolia dùng node có sẵn của team: `37.27.127.216` (cổng 8545, 8546, 3500).

## 1. Server

- 8 vCPU, 16 GB RAM, 20 GB SSD
- Mở cổng vào **9651** (p2p Avalanche). Cổng 9650, 18080, 19090 chỉ dùng nội bộ.

## 2. Key (lấy từ secret manager, không commit)

| Key | Dùng cho | Cần nạp |
|---|---|---|
| Relayer key | `icm-relayer` | ~0.1 Sepolia ETH + ~0.1 Fuji AVAX |
| Updater key | `lc-updater` | ~0.1 Fuji AVAX |
| Registry-updater key | `icm-relayer` (bắt buộc khác relayer key) | không cần |

Mỗi key chỉ cho **một** tiến trình dùng, nếu không sẽ trùng nonce.

## 3. Build (Go 1.26)

```sh
# avalanchego: tải bản release v1.15.2
curl -LO https://github.com/ava-labs/avalanchego/releases/download/v1.15.2/avalanchego-linux-amd64-v1.15.2.tar.gz

# updater (đúng commit dd95dd5, khớp với proving key)
git clone git@github.com:decentrio/avax-light-client.git && cd avax-light-client
git checkout dd95dd5 && go build -o /opt/bridge/bin/updater ./cmd/updater

# relayer
git clone git@github.com:decentrio/icm-services.git && cd icm-services
git checkout hieu/sepolia-fuji-demo && go build -o /opt/bridge/bin/icm-relayer ./relayer/main
```

Copy thêm thư mục proving key `avax-light-client/keys/` (3,7 GB, lấy từ máy dev) vào `/var/lib/lc-updater/keys`. **Không tạo lại key.**

## 4. Cấu hình

**Relayer:** không cần biến môi trường. Mọi thứ nằm trong một file JSON:

```sh
cp deploy/relayer.example.json /etc/bridge/relayer.json
chmod 600 /etc/bridge/relayer.json
# điền: <SEPOLIA_EL>, <SEPOLIA_CL> = 37.27.127.216
#       <RELAYER_KEY>, <RELAYER_ADDRESS>, <REGISTRY_UPDATER_KEY>
```

File này chứa key, nên chỉ user chạy service được đọc. Trên Docker/K8s, mount file này từ secret.

**Updater:** cần biến môi trường:

```sh
# /etc/bridge/updater.env  (chmod 600)
RELAYER_KEY=0x<updater key>
```

`GOMEMLIMIT=6GiB` đã có sẵn trong unit.

## 5. Chạy

Copy 3 file `deploy/*.service` vào `/etc/systemd/system/`, rồi chạy **theo thứ tự**:

```sh
systemctl enable --now avalanchego-fuji
# đợi ~10 phút đến khi trả về "isBootstrapped":true
curl -s -X POST -H 'content-type:application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"info.isBootstrapped","params":{"chain":"P"}}' \
  http://127.0.0.1:9650/ext/info

systemctl enable --now lc-updater     # nạp key mất ~3 phút
systemctl enable --now icm-relayer
```

Cả 3 unit đều có `Restart=always`. Relayer **tự thoát khi mất mạng**, nên phải để systemd (hoặc Docker/K8s) bật lại.

Khi server đã chạy thì **tắt các tiến trình trên laptop**, vì cùng key.

## 6. Kiểm tra

```sh
curl -s localhost:18080/health        # relayer: "status":"up"
journalctl -u lc-updater -f           # updater: "tick done ... latest finalized slot N" mỗi ~12 giây
```

Thử chuyển 1 EXMP Sepolia → Fuji (mất ~15–20 phút) và 0,5 EXMP.e Fuji → Sepolia (mất ~10–30 giây), bằng trang Bridge trên UI. Log relayer phải có dòng `Delivered message to destination chain`.

**Cảnh báo khi:**
- health khác `up`;
- message Sepolia → Fuji chờ quá 30 phút, hoặc Fuji → Sepolia chờ quá 2 phút;
- relayer key còn < 0,01 ETH hoặc < 0,05 AVAX, updater key còn < 0,02 AVAX.

## Chuyển server đang chạy sang light client v2 (mạch 2,86M)

Chỉ dành cho server đã chạy bản cũ (light client `0x4fc0…5e5c`). Dev đã deploy sẵn light client, adapter và messenger v2; cặp EXMP ⇄ EXMP.e giữ nguyên địa chỉ và số dư.

1. **Chuẩn bị** (hệ thống cũ vẫn chạy):
   - Build updater mới: `git checkout dd95dd5 && go build -o /opt/bridge/bin/updater-v2 ./cmd/updater` trong `avax-light-client`.
   - Copy proving key mới (3,7 GB, dev gửi) vào `/var/lib/lc-updater/keys-v2`.
   - Đặt `relayer.json` mới (dev gửi, chỉ khác 3 địa chỉ) ở `/etc/bridge/relayer-v2.json`, mode 600.
2. **Chờ hết giao dịch đang đi dở:** trên trang Monitor, mọi dòng đều `Delivered`.
3. **Dừng bản cũ:** `systemctl stop icm-relayer lc-updater`.
4. **Báo dev đổi route:** dev chuyển cặp EXMP sang messenger mới (4 giao dịch từ ví admin, khoảng 1 phút).
5. **Chạy bản mới:**
   ```sh
   mv /etc/bridge/relayer-v2.json /etc/bridge/relayer.json
   # trong /etc/systemd/system/lc-updater.service:
   #   /opt/bridge/bin/updater          → /opt/bridge/bin/updater-v2
   #   -contract 0x4fc0…  → -contract 0xb72fefdfde361319ebb54bb2f9fef5fa70b42efe
   #   -adapter  0x59d6…  → -adapter  0xf76523ba1dc38a930c63a0e64ca5efd853f52413
   #   -keys /var/lib/lc-updater/keys → -keys /var/lib/lc-updater/keys-v2
   systemctl daemon-reload && systemctl start lc-updater icm-relayer
   ```
   Updater key giữ nguyên (`0x51Fd…`, đã được cấp quyền trên light client v2).
6. **Deploy UI** nhánh `feat/fuji-sepolia-monitor` (commit `1c789b1` trở đi).
7. **Kiểm tra:** log updater có `stepping on demand for messages from 0xF765…`, rồi chuyển thử 1 EXMP mỗi chiều.

**Quay lại bản cũ nếu lỗi:** dev đổi route cặp EXMP về messenger cũ, infra khôi phục `relayer.json` và unit cũ.

## Giới hạn (chỉ dùng cho testnet)

- Bộ validator Fuji lưu trên Sepolia **cố định**, dùng được đến khoảng **03/2027**. Sau đó cần dev deploy lại registry.
- Proving key là bản dev, chưa qua ceremony. **Không dùng cho mainnet.**

Chi tiết đầy đủ bằng tiếng Anh: [DEPLOY.md](DEPLOY.md).

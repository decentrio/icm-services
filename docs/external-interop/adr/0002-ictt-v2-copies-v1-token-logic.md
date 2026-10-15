# ICTT v2 copies v1's token logic instead of sharing it

ICTT v2's TokenHome and TokenRemote are copies of v1's, with only the messenger lookup, the source of the blockchain ID and initialization changed. About 2,300 lines of token logic are duplicated, so a fix to that logic has to be made in both. We considered moving the token logic into a core that v1 and v2 both inherit. We rejected it because it means editing audited, deployed v1 contracts and putting their storage layout at risk, while new deployments move to v2 anyway. Revisit this if v1 stays under active development.

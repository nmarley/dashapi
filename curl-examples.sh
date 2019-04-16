#! /bin/sh

# Examples of using cURL to access the API. Replace URL here
# w/the actual URL of service deployment for actual usage.

# TODO:
# example of posting invalid data
#
# curl --header "Content-Type: application/json" --data '{"addr": "7", "msg": "hi", "sig": "yo"}' http://localhost:7001/vote

# TODO:
# example of posting valid data
#
# curl --header "Content-Type: application/json" --data '{"addr": "yMtMWAjPhUquwKtdG4wzj9Cpn4asQkLV8F", "msg": "dte2019-efigaro|lcole|sfigaro|cchere", "sig": "HwFI6cUwJMLhB2koK5BcBxFQgLHCrrhKg+28TKO7B3eVRH33B987NBrlpo80xETBPF7xjHs7AUflVqjB/MpLetE="}' http://localhost:7001/vote

# TODO:
# example listing valid proposals
#
# "valid" votes are those which are not superceded by any newer vote for the
# same MNO collateral address
#
# note: JWT_TOKEN must be set to a valid, signed token
#
# curl --silent --header "Authorization: Bearer $JWT_TOKEN" http://localhost:7001/validVotes

# TODO:
# example listing all proposals
#
# note: JWT_TOKEN must be set to a valid, signed token
#
# curl --silent --header "Authorization: Bearer $JWT_TOKEN" http://localhost:7001/allVotes

# TODO:
# * triggers

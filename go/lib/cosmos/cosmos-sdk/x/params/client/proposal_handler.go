package client

import (
	govclient "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/params/client/cli"
	"github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)

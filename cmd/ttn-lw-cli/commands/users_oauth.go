// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/v3/cmd/internal/io"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func getUserAndClientID(flagSet *pflag.FlagSet, args []string) (*ttnpb.UserIdentifiers, *ttnpb.ClientIdentifiers) {
	userID, _ := flagSet.GetString("user-id")
	clientID, _ := flagSet.GetString("client-id")
	switch len(args) {
	case 0:
	case 1:
		logger.Warn("Only single ID found in arguments, not considering arguments")
	case 2:
		userID = args[0]
		clientID = args[1]
	default:
		logger.Warn("Multiple IDs found in arguments, considering the first")
		userID = args[0]
		clientID = args[1]
	}
	switch {
	case userID != "" && clientID != "":
		return &ttnpb.UserIdentifiers{UserId: userID}, &ttnpb.ClientIdentifiers{ClientId: clientID}
	case userID != "":
		return &ttnpb.UserIdentifiers{UserId: userID}, nil
	case clientID != "":
		return nil, &ttnpb.ClientIdentifiers{ClientId: clientID}
	}
	return nil, nil
}

var errNoTokenID = errors.DefineInvalidArgument("no_token_id", "no token ID set")

var (
	oauthCommand = &cobra.Command{
		Use:   "oauth",
		Short: "Manage OAuth authorizations and access tokens",
	}
	oauthAuthorizationsCommand = &cobra.Command{
		Use:   "authorizations",
		Short: "Manage OAuth authorizations",
	}
	oauthAuthorizationsListCommand = &cobra.Command{
		Use:     "list [user-id]",
		Aliases: []string{"ls"},
		Short:   "List OAuth authorizations",
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID := getUserID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			limit, page, opt, getTotal := withPagination(cmd.Flags())
			res, err := ttnpb.NewOAuthAuthorizationRegistryClient(is).List(ctx, &ttnpb.ListOAuthClientAuthorizationsRequest{
				UserIdentifiers: *usrID, Limit: limit, Page: page, Order: getOrder(cmd.Flags()),
			}, opt)
			if err != nil {
				return err
			}
			getTotal()

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	oauthAuthorizationsDeleteCommand = &cobra.Command{
		Use:     "delete [user-id] [client-id]",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete an OAuth authorization and all access tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID, cliID := getUserAndClientID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}
			if cliID == nil {
				return errNoClientID.New()
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}

			res, err := ttnpb.NewOAuthAuthorizationRegistryClient(is).ListTokens(ctx, &ttnpb.ListOAuthAccessTokensRequest{
				UserIds:   *usrID,
				ClientIds: *cliID,
			})
			if err != nil {
				return err
			}

			for _, token := range res.Tokens {
				_, err = ttnpb.NewOAuthAuthorizationRegistryClient(is).DeleteToken(ctx, &ttnpb.OAuthAccessTokenIdentifiers{
					UserIds:   *usrID,
					ClientIds: *cliID,
					Id:        token.Id,
				})
				if err != nil {
					return err
				}
			}

			_, err = ttnpb.NewOAuthAuthorizationRegistryClient(is).Delete(ctx, &ttnpb.OAuthClientAuthorizationIdentifiers{
				UserIds:   *usrID,
				ClientIds: *cliID,
			})

			return err
		},
	}
	oauthAccessTokensCommand = &cobra.Command{
		Use:     "access-tokens",
		Aliases: []string{"tokens"},
		Short:   "Manage OAuth tokens",
	}
	oauthAccessTokensListCommand = &cobra.Command{
		Use:     "list [user-id] [client-id]",
		Aliases: []string{"ls"},
		Short:   "List OAuth access tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID, cliID := getUserAndClientID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}
			if cliID == nil {
				return errNoClientID.New()
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			limit, page, opt, getTotal := withPagination(cmd.Flags())
			res, err := ttnpb.NewOAuthAuthorizationRegistryClient(is).ListTokens(ctx, &ttnpb.ListOAuthAccessTokensRequest{
				UserIds:   *usrID,
				ClientIds: *cliID,
				Limit:     limit,
				Page:      page,
				Order:     getOrder(cmd.Flags()),
			}, opt)
			if err != nil {
				return err
			}
			getTotal()

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	oauthAccessTokensDeleteCommand = &cobra.Command{
		Use:     "delete [user-id] [client-id]",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete an OAuth access token",
		RunE: func(cmd *cobra.Command, args []string) error {
			usrID, cliID := getUserAndClientID(cmd.Flags(), args)
			if usrID == nil {
				return errNoUserID.New()
			}
			if cliID == nil {
				return errNoClientID.New()
			}
			tokenID, _ := cmd.Flags().GetString("token-id")
			if tokenID == "" {
				return errNoTokenID.New()
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}

			_, err = ttnpb.NewOAuthAuthorizationRegistryClient(is).DeleteToken(ctx, &ttnpb.OAuthAccessTokenIdentifiers{
				UserIds:   *usrID,
				ClientIds: *cliID,
				Id:        tokenID,
			})

			return err
		},
	}
)

func init() {
	oauthAuthorizationsListCommand.Flags().AddFlagSet(userIDFlags())
	oauthAuthorizationsListCommand.Flags().AddFlagSet(paginationFlags())
	oauthAuthorizationsCommand.AddCommand(oauthAuthorizationsListCommand)
	oauthAuthorizationsDeleteCommand.Flags().AddFlagSet(userIDFlags())
	oauthAuthorizationsDeleteCommand.Flags().AddFlagSet(clientIDFlags())
	oauthAuthorizationsCommand.AddCommand(oauthAuthorizationsDeleteCommand)
	oauthCommand.AddCommand(oauthAuthorizationsCommand)
	oauthAccessTokensListCommand.Flags().AddFlagSet(userIDFlags())
	oauthAccessTokensListCommand.Flags().AddFlagSet(clientIDFlags())
	oauthAccessTokensListCommand.Flags().AddFlagSet(paginationFlags())
	oauthAccessTokensCommand.AddCommand(oauthAccessTokensListCommand)
	oauthAccessTokensDeleteCommand.Flags().AddFlagSet(userIDFlags())
	oauthAccessTokensDeleteCommand.Flags().AddFlagSet(clientIDFlags())
	oauthAccessTokensDeleteCommand.Flags().String("token-id", "", "")
	oauthAccessTokensCommand.AddCommand(oauthAccessTokensDeleteCommand)
	oauthCommand.AddCommand(oauthAccessTokensCommand)
	usersCommand.AddCommand(oauthCommand)
}

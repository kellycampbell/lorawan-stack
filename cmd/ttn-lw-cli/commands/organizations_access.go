// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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
	"strings"

	"github.com/spf13/cobra"
	"go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/internal/io"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

var (
	organizationRights = &cobra.Command{
		Use:   "rights",
		Short: "List the rights to an organization",
		RunE: func(cmd *cobra.Command, args []string) error {
			orgID := getOrganizationID(cmd.Flags(), args)
			if orgID == nil {
				return errNoOrganizationID
			}

			is, err := api.Dial(ctx, config.IdentityServerAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewOrganizationAccessClient(is).ListRights(ctx, orgID)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.Format, res.Rights)
		},
	}
	organizationCollaborators = &cobra.Command{
		Use:     "collaborators",
		Aliases: []string{"collaborator", "members", "member"},
		Short:   "Manage organization collaborators",
	}
	organizationCollaboratorsList = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List organization collaborators",
		RunE: func(cmd *cobra.Command, args []string) error {
			orgID := getOrganizationID(cmd.Flags(), args)
			if orgID == nil {
				return errNoOrganizationID
			}

			is, err := api.Dial(ctx, config.IdentityServerAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewOrganizationAccessClient(is).ListCollaborators(ctx, orgID)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.Format, res.Collaborators)
		},
	}
	organizationCollaboratorsSet = &cobra.Command{
		Use:   "set",
		Short: "Set an organization collaborator",
		RunE: func(cmd *cobra.Command, args []string) error {
			orgID := getOrganizationID(cmd.Flags(), nil)
			if orgID == nil {
				return errNoOrganizationID
			}
			collaborator := getCollaborator(cmd.Flags())
			if collaborator == nil {
				return errNoCollaborator
			}
			rights := getRights(cmd.Flags())
			if len(rights) == 0 {
				logger.Info("No rights selected, will remove collaborator")
			}

			is, err := api.Dial(ctx, config.IdentityServerAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewOrganizationAccessClient(is).SetCollaborator(ctx, &ttnpb.SetOrganizationCollaboratorRequest{
				OrganizationIdentifiers: *orgID,
				Collaborator: ttnpb.Collaborator{
					OrganizationOrUserIdentifiers: *collaborator,
					Rights:                        rights,
				},
			})
			if err != nil {
				return err
			}

			return nil
		},
	}
	organizationAPIKeys = &cobra.Command{
		Use:     "api-keys",
		Aliases: []string{"api-key"},
		Short:   "Manage organization API keys",
	}
	organizationAPIKeysList = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List organization API keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			orgID := getOrganizationID(cmd.Flags(), args)
			if orgID == nil {
				return errNoOrganizationID
			}

			is, err := api.Dial(ctx, config.IdentityServerAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewOrganizationAccessClient(is).ListAPIKeys(ctx, orgID)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.Format, res.APIKeys)
		},
	}
	organizationAPIKeysCreate = &cobra.Command{
		Use:     "create",
		Aliases: []string{"add", "generate"},
		Short:   "Create an organization API key",
		RunE: func(cmd *cobra.Command, args []string) error {
			orgID := getOrganizationID(cmd.Flags(), nil)
			if orgID == nil {
				return errNoOrganizationID
			}
			name, _ := cmd.Flags().GetString("name")

			rights := getRights(cmd.Flags())
			if len(rights) == 0 {
				logger.Info("No rights selected, won't create API key")
				return nil
			}

			is, err := api.Dial(ctx, config.IdentityServerAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewOrganizationAccessClient(is).CreateAPIKey(ctx, &ttnpb.CreateOrganizationAPIKeyRequest{
				OrganizationIdentifiers: *orgID,
				Name:                    name,
				Rights:                  rights,
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.Format, res)
		},
	}
	organizationAPIKeysUpdate = &cobra.Command{
		Use:     "update",
		Aliases: []string{"set"},
		Short:   "Update an organization API key",
		RunE: func(cmd *cobra.Command, args []string) error {
			id := getAPIKeyID(cmd.Flags(), args)
			if id == "" {
				return errNoAPIKeyID
			}
			orgID := getOrganizationID(cmd.Flags(), nil)
			if orgID == nil {
				return errNoOrganizationID
			}
			name, _ := cmd.Flags().GetString("name")

			rights := getRights(cmd.Flags())
			if len(rights) == 0 {
				logger.Info("No rights selected, will remove API key")
			}

			is, err := api.Dial(ctx, config.IdentityServerAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewOrganizationAccessClient(is).UpdateAPIKey(ctx, &ttnpb.UpdateOrganizationAPIKeyRequest{
				OrganizationIdentifiers: *orgID,
				APIKey: ttnpb.APIKey{
					ID:     id,
					Name:   name,
					Rights: rights,
				},
			})
			if err != nil {
				return err
			}

			return nil
		},
	}
)

var organizationRightsFlags = rightsFlags(func(flag string) bool {
	return strings.HasPrefix(flag, "right-organization")
})

func init() {
	organizationRights.Flags().AddFlagSet(organizationIDFlags())
	organizationsCommand.AddCommand(organizationRights)

	organizationCollaborators.AddCommand(organizationCollaboratorsList)
	organizationCollaboratorsSet.Flags().AddFlagSet(collaboratorFlags())
	organizationCollaboratorsSet.Flags().AddFlagSet(organizationRightsFlags)
	organizationCollaborators.AddCommand(organizationCollaboratorsSet)
	organizationCollaborators.PersistentFlags().AddFlagSet(organizationIDFlags())
	organizationsCommand.AddCommand(organizationCollaborators)

	organizationAPIKeys.AddCommand(organizationAPIKeysList)
	organizationAPIKeysCreate.Flags().String("name", "", "")
	organizationAPIKeysCreate.Flags().AddFlagSet(organizationRightsFlags)
	organizationAPIKeys.AddCommand(organizationAPIKeysCreate)
	organizationAPIKeysUpdate.Flags().String("api-key-id", "", "")
	organizationAPIKeysUpdate.Flags().String("name", "", "")
	organizationAPIKeysUpdate.Flags().AddFlagSet(organizationRightsFlags)
	organizationAPIKeys.AddCommand(organizationAPIKeysUpdate)
	organizationAPIKeys.PersistentFlags().AddFlagSet(organizationIDFlags())
	organizationsCommand.AddCommand(organizationAPIKeys)
}

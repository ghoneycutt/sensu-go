import React from "react";
import PropTypes from "prop-types";
import partition from "lodash/partition";
import { map } from "lodash";
import { compose } from "recompose";
import { withStyles } from "material-ui/styles";
import { createFragmentContainer, graphql } from "react-relay";

import Menu, { MenuItem } from "material-ui/Menu";
import { ListItemIcon, ListItemText } from "material-ui/List";
import Divider from "material-ui/Divider";
import OrganizationIcon from "./OrganizationIcon";
import EnvironmentSymbol from "./EnvironmentSymbol";
import { withNamespace, namespaceShape } from "./NamespaceLink";

const styles = () => ({
  envIcon: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    width: 24,
    height: 24,
  },
});

class NamespaceSelectorMenu extends React.Component {
  static propTypes = {
    classes: PropTypes.object.isRequired,
    currentNamespace: namespaceShape.isRequired,
    viewer: PropTypes.objectOf(PropTypes.any).isRequired,
  };

  render() {
    const { viewer, currentNamespace, classes, ...props } = this.props;
    const navigateTo = (organization, environment) => () =>
      window.location.assign(`/${organization}/${environment}/`);

    const partitionedOrganizations = partition(
      viewer.organizations,
      organization => organization.name === currentNamespace.organization,
    );
    const currentOrganization = partitionedOrganizations[0][0];
    const otherOrganizations = partitionedOrganizations[1];

    return (
      <Menu {...props}>
        {currentOrganization.environments.map(environment => (
          <MenuItem
            key={environment.name}
            onClick={navigateTo(currentOrganization.name, environment.name)}
          >
            <ListItemIcon>
              <div className={classes.envIcon}>
                <EnvironmentSymbol environment={environment} size={12} />
              </div>
            </ListItemIcon>
            <ListItemText inset primary={environment.name} />
          </MenuItem>
        ))}
        <Divider />
        {map(otherOrganizations, (organization, i) => [
          organization.environments.map(environment => (
            <MenuItem
              key={environment.name}
              onClick={navigateTo(organization.name, environment.name)}
            >
              <ListItemIcon>
                <OrganizationIcon organization={organization} size={24}>
                  <EnvironmentSymbol environment={environment} size={24 / 3} />
                </OrganizationIcon>
              </ListItemIcon>
              <ListItemText
                inset
                primary={organization.name}
                secondary={environment.name}
              />
            </MenuItem>
          )),
          i + 1 < otherOrganizations.length ? <Divider /> : null,
        ])}
      </Menu>
    );
  }
}

export default createFragmentContainer(
  compose(withStyles(styles), withNamespace)(NamespaceSelectorMenu),
  graphql`
    fragment NamespaceSelectorMenu_viewer on Viewer {
      organizations {
        name
        ...OrganizationIcon_organization
        environments {
          name
          ...EnvironmentSymbol_environment
        }
      }
    }
  `,
);

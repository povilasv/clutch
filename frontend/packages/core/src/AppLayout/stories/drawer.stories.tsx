import * as React from "react";
import { MemoryRouter } from "react-router";
import { Grid } from "@material-ui/core";
import type { Meta } from "@storybook/react";

import { ApplicationContext } from "../../Contexts/app-context";
import Drawer from "../drawer";
import Header from "../header";

export default {
  title: "Core/AppLayout/Drawer",
  component: Drawer,
  decorators: [
    StoryFn => (
      <MemoryRouter>
        <StoryFn />
      </MemoryRouter>
    ),
    StoryFn => {
      const value = React.useMemo(
        () => ({
          workflows: [
            {
              developer: { name: "Lyft", contactUrl: "mailto:hello@clutch.sh" },
              displayName: "EC2",
              group: "AWS",
              path: "ec2",
              routes: [
                {
                  component: () => null,
                  componentProps: { resolverType: "clutch.aws.ec2.v1.Instance" },
                  description: "Terminate an EC2 instance.",
                  displayName: "Terminate Instance",
                  path: "instance/terminate",
                  requiredConfigProps: ["resolverType"],
                  trending: true,
                },
                {
                  component: () => null,
                  componentProps: { resolverType: "clutch.aws.ec2.v1.AutoscalingGroup" },
                  description: "Resize an autoscaling group.",
                  displayName: "Resize Autoscaling Group",
                  path: "asg/resize",
                  requiredConfigProps: ["resolverType"],
                },
              ],
            },
          ],
        }),
        []
      );
      return (
        <ApplicationContext.Provider value={value}>
          <StoryFn />
        </ApplicationContext.Provider>
      );
    },
  ],
  parameters: {
    layout: "fullscreen",
  },
} as Meta;

export function Primary() {
  return <Drawer />;
}

export function WithHeader() {
  return (
    <Grid direction="column">
      <Header />
      <Drawer />
    </Grid>
  );
}

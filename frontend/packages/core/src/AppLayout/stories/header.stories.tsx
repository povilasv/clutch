import * as React from "react";
import { MemoryRouter } from "react-router";
import type { Meta } from "@storybook/react";

import { ApplicationContext } from "../../Contexts/app-context";
import Header from "../header";

export default {
  title: "Core/AppLayout/Header",
  component: Header,
  decorators: [
    () => (
      <MemoryRouter>
        <Header />
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
                  component: () => <div>Terminate Instance</div>,
                  componentProps: { resolverType: "clutch.aws.ec2.v1.Instance" },
                  description: "Terminate an EC2 instance.",
                  displayName: "Terminate Instance",
                  path: "instance/terminate",
                  requiredConfigProps: ["resolverType"],
                  trending: true,
                },
                {
                  component: () => <div>Resize ASG</div>,
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

export const Primary: React.FC<{}> = () => <Header />;

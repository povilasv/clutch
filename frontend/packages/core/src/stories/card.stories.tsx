import * as React from "react";
import type { Meta } from "@storybook/react";

import type { CardProps } from "../card";
import { Card, CardContent } from "../card";

export default {
  title: "Core/Card/Basic",
  component: Card,
} as Meta;

function Template(props: CardProps) {
  return <Card {...props} />;
}

export const Primary = Template.bind({});
Primary.args = {
  children: <CardContent>Hello world!</CardContent>,
};

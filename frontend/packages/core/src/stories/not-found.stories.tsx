import React from "react";
import type { Meta } from "@storybook/react";

import NotFound from "../not-found";

export default {
  title: "Core/NotFound",
  component: NotFound,
} as Meta;

function Template() {
  return <NotFound />;
}
export const Default = Template.bind({});

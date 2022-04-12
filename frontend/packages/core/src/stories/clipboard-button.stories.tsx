import React from "react";
import type { Meta } from "@storybook/react";

import type { ClipboardButtonProps } from "../button";
import { ClipboardButton } from "../button";

export default {
  title: "Core/Buttons/Clipboard Button",
  component: ClipboardButton,
} as Meta;

function Template(props: ClipboardButtonProps) {
  return <ClipboardButton {...props} />;
}

export const Default = Template.bind({});
Default.args = {
  text: "https://clutch.sh",
};

import * as React from "react";
import type { Meta } from "@storybook/react";

import Paper from "../paper";

export default {
  title: "Core/Paper",
  component: Paper,
} as Meta;

export function Primary() {
  return (
    <Paper>
      <div>Some text in paper</div>
    </Paper>
  );
}

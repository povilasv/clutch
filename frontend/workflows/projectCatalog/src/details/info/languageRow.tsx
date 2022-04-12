import React from "react";
import { Grid, Typography } from "@clutch-sh/core";

import LanguageIcon from "../../helpers/language-icon";

function LanguageRow({ languages }: { languages: string[] }) {
  return (
    <>
      <Grid item>
        <Typography variant="body2">Language</Typography>
      </Grid>
      {languages.map(language => (
        <Grid item>
          <LanguageIcon language={language} />
        </Grid>
      ))}
    </>
  );
}

export default LanguageRow;

import GroupIcon from "@mui/icons-material/Groups";
import { Button, Grid } from "@mui/material";
import Link from "next/link";
import { Page } from "../components/Page";
import { TeamLogo } from "../components/TeamLogo";

export default function Home() {
  return (
    <Page>
      <Grid
        container
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Grid item>
          <TeamLogo />
        </Grid>
        <Grid item>
          <Button
            component={Link}
            href="/players"
            variant="contained"
            startIcon={<GroupIcon />}
          >
            Choose players
          </Button>
        </Grid>
      </Grid>
    </Page>
  );
}

import GroupIcon from "@mui/icons-material/Groups";
import { Button, Grid } from "@mui/material";
import Link from "next/link";
import { Page } from "../components/Page";

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
        <Grid item></Grid>
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

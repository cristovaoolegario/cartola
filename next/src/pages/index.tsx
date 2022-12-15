import GroupIcon from "@mui/icons-material/Groups";
import { Button, Divider, Grid, styled } from "@mui/material";
import { NextPage } from "next";
import Link from "next/link";
import { Label } from "../components/Label";
import { Page } from "../components/Page";
import { Section } from "../components/Section";
import { TeamLogo } from "../components/TeamLogo";

const BudgetContainer = styled(Section)(({ theme }) => ({
  width: "888px",
  height: "300px",
  marginTop: theme.spacing(8),
  display: "flex",
  alignItems: "center",
}));

const HomePage: NextPage = () => {
  return (
    <Page>
      <Grid
        container
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          gap: (theme) => theme.spacing(3),
        }}
      >
        <Grid item>
          <TeamLogo
            sx={{ position: "absolute", left: 0, right: 0, m: "auto" }}
          />
          <BudgetContainer>
            <Grid container>
              <Grid
                item
                xs={5}
                sx={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                }}
              >
                <Label>Latest score</Label>
                <Label>99.84</Label>
              </Grid>
              <Grid
                item
                xs={2}
                sx={{ display: "flex", justifyContent: "center" }}
              >
                <Divider orientation="vertical" sx={{ height: "auto" }} />
              </Grid>
              <Grid
                item
                xs={5}
                sx={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "center",
                }}
              >
                <Label>Patrimony</Label>
                <Label>300</Label>
              </Grid>
            </Grid>
          </BudgetContainer>
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
};

export default HomePage;

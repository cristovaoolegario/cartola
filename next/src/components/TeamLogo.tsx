import { Box, BoxProps, Typography } from "@mui/material";
import Image from "next/image";
import { Label } from "./Label";

export type TeamLogoProps = BoxProps;

export const TeamLogo = (props: TeamLogoProps) => {
  return (
    <Box
      {...props}
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        ...props.sx,
      }}
    >
      <Image
        src="/img/my-team-logo.svg"
        width={84.5}
        height={88.5}
        alt="My Team FC"
        priority={true}
      />
      <Label>My Team FC</Label>
    </Box>
  );
};

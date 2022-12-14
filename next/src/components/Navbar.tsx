import { AppBar, Box, Button, Toolbar } from "@mui/material";
import Image from "next/image";
import Link, { LinkProps } from "next/link";
import { PropsWithChildren } from "react";

export type NavbarItemProps = LinkProps;

export const NavbarItem = (props: PropsWithChildren<NavbarItemProps>) => {
  //@ts-expect-error
  return <Button component={Link} sx={{ color: "white" }} {...props} />;
};

export const Navbar = () => {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static" sx={{ background: "none", boxShadow: "none" }}>
        <Toolbar>
          <Image
            src="/img/logo.png"
            width={315}
            height={58}
            alt="logo"
            priority={true}
          />
          <Box sx={{ flexGrow: 1 }}>
            <NavbarItem href="/">Home</NavbarItem>
            <NavbarItem href="/players">Players</NavbarItem>
            <NavbarItem href="/matches">Matches</NavbarItem>
          </Box>
        </Toolbar>
      </AppBar>
    </Box>
  );
};

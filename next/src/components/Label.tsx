import { Typography, TypographyProps } from "@mui/material";

export type Labelprops = TypographyProps;

export const Label = (props: Labelprops) => {
  return <Typography variant="h6" component="span" {...props} />;
};

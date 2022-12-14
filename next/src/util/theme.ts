import { createTheme } from "@mui/material";
import { Roboto } from "@next/font/google";

export const roboto = Roboto({
  weight: ["300", "400", "500", "700"],
  subsets: ["latin"],
  display: "swap",
  fallback: ["Helvetica", "Arial", "sans-serif"],
});

const theme = createTheme({
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        "html, body, body>div": {
          padding: 0,
          margin: 0,
          width: "100%",
          height: "100%",
        },
        body: {
          background: "url(/img/background.png) no-repeat center center fixed",
          WebkitBackgroundSize: "cover",
          OBackgroundSize: "cover",
          BackgroundSize: "cover",
          MozBackgroundSize: "cover",
        },
      },
    },
  },
});

export default theme;

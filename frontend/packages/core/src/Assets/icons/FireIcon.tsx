import * as React from "react";

import StyledSvg from "./helpers";

const FireIcon = (props: React.SVGProps<SVGSVGElement>) => (
  <StyledSvg
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    {...props}
  >
    <path
      d="M19.48 12.35C17.91 8.27 12.32 8.05 13.67 2.12C13.77 1.68 13.3 1.34 12.92 1.57C9.29 3.71 6.68 8 8.87 13.62C9.05 14.08 8.51 14.51 8.12 14.21C6.31 12.84 6.12 10.87 6.28 9.46C6.34 8.94 5.66 8.69 5.37 9.12C4.69 10.16 4 11.84 4 14.37C4.38 19.97 9.11 21.69 10.81 21.91C13.24 22.22 15.87 21.77 17.76 20.04C19.84 18.11 20.6 15.03 19.48 12.35ZM10.2 17.38C11.64 17.03 12.38 15.99 12.58 15.07C12.91 13.64 11.62 12.24 12.49 9.98C12.82 11.85 15.76 13.02 15.76 15.06C15.84 17.59 13.1 19.76 10.2 17.38Z"
      fill="#F59E0B"
    />
  </StyledSvg>
);

export default FireIcon;

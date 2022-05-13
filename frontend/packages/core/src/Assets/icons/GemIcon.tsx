import * as React from "react";

import type { SVGProps } from "../global";
import { StyledSVG } from "../global";

const GemIcon = ({ size, ...props }: SVGProps) => (
  <StyledSVG
    size={size}
    viewBox="0 0 24 24"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    {...props}
  >
    <path
      d="M15.4051 6.65088L9.9846 17.5073L7.8669 13.266L4.56425 6.65088H15.4051Z"
      fill="#31D2F7"
    />
    <path
      d="M3.78246 5.66151L3.5331 5.61156L0.609638 4.83472L4.20339 1.33472L5.95278 1.99487L4.02722 5.24796L3.78246 5.66151Z"
      fill="#31D2F7"
    />
    <path
      d="M15.9994 5.66151L16.2487 5.61156L19.1722 4.83472L15.5785 1.33472L13.8291 1.99487L15.7546 5.24796L15.9994 5.66151Z"
      fill="#31D2F7"
    />
    <path
      d="M4.79267 5.73417L6.90329 2.16858H13.066L13.4079 2.74651L15.1764 5.73417H4.79267Z"
      fill="#31D2F7"
    />
    <path
      d="M3.1263 6.51818L3.61831 6.59077L3.64491 6.65082L8.82842 16.7573L0.361299 5.68219L3.1263 6.51818Z"
      fill="#31D2F7"
    />
    <path
      d="M16.8737 6.51818L16.3817 6.59077L16.3551 6.65082L11.1716 16.7573L19.6387 5.68219L16.8737 6.51818Z"
      fill="#31D2F7"
    />
    <path
      d="M6.61455 1.29803L5.65598 1.11412L4.87768 0.964853L7.04265 0.509777L12.8231 0.492554L15.0851 0.966002L14.5342 1.07164L13.3548 1.29803H6.61455Z"
      fill="#31D2F7"
    />
  </StyledSVG>
);

export default GemIcon;

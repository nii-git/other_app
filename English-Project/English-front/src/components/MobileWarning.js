import { BrowserView, MobileView, isBrowser, isMobile } from 'react-device-detect';
import React from "react";



export const MobileWarning =() => {
  return (
    <>
    <div>
    <MobileView>
    <h1 style={{color:"red"}}>スマホでの閲覧は非推奨です!</h1>
    </MobileView>
    </div>
    </>
  );
}
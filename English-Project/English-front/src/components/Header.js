// TopPage.js

import React, { useEffect, useState } from "react";
import {
    HeaderNavigation,
    ALIGN,
    StyledNavigationList,
    StyledNavigationItem
  } from "baseui/header-navigation";
  import { StyledLink } from "baseui/link";
  import { Button } from "baseui/button";
  import logoImg from "./../img/sample-logo.png"
  import { BrowserView, MobileView, isBrowser, isMobile } from 'react-device-detect';


export const Header = () => {
  return (
    <>
    <HeaderNavigation>
        <StyledNavigationList $align={ALIGN.left}>
        <div> <a href="/"><img src={logoImg} class="logoimg"></img></a></div>
        {/* <StyledNavigationItem>TITLE</StyledNavigationItem> */}
        </StyledNavigationList>
        <StyledNavigationList $align={ALIGN.center} />
        <StyledNavigationList $align={ALIGN.right}>
        <StyledNavigationItem>
            <StyledLink href="/about">
            About
            </StyledLink>
        </StyledNavigationItem>
        <StyledNavigationItem>
            <StyledLink href="/providers">
            Websites
            </StyledLink>
        </StyledNavigationItem>
        </StyledNavigationList>
        <StyledNavigationList $align={ALIGN.right}>
        <StyledNavigationItem>
            <Button>Search</Button>
        </StyledNavigationItem>
        </StyledNavigationList>
    </HeaderNavigation>
    <div>
    <MobileView>
    <h1 style={{color:"red"}}>スマホでの閲覧は非推奨です!</h1>
    </MobileView>
    </div>
    </>
  );
}
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
  import logoImg from "./../img/logo.png"
  import { BrowserView, MobileView, isBrowser, isMobile } from 'react-device-detect';
  import { Drawer } from "baseui/drawer";
  import { StatefulMenu } from "baseui/menu";



  function mobileSideBar() {

    return (
      <Drawer
        isOpen={isOpen}
        autoFocus
        onClose={() => setIsOpen(false)}
      >
        <div>drawer content</div>
      </Drawer>
    );
  }

export const Header = () => {
  const [isOpen, setIsOpen] = React.useState(false);

  useEffect(() => {},[]);

  const toggleDrawer = () => {
    setIsOpen(!isOpen);
};

  return (
    <>
    <BrowserView>
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
          {/* 実装次第Searchに変えます */}
          <StyledLink href="/">
            <Button>TopPage</Button>
          </StyledLink>
        </StyledNavigationItem>
        </StyledNavigationList>
    </HeaderNavigation>
    </BrowserView>
    <MobileView>
      <Drawer
          isOpen={isOpen}
          autoFocus
          onClose={() => setIsOpen(false)}
          overrides={{
            DrawerContainer: {
              style: ({ $theme }) => ({
                outline: `${$theme.colors.white} solid`,
                backgroundColor: "transparent"
              })
            }
          }}
      >
      <StatefulMenu
        items={[
          {label: 'TOP', href: '/'},
          {label: 'About', href: '/about'},
          {label: 'Websites', href: '/providers'},
        ]}
      />
      </Drawer>
      <HeaderNavigation>
          <StyledNavigationList $align={ALIGN.left}>
              <div> <a href="/"><img src={logoImg} class="logoimgMobile"></img></a></div>
          </StyledNavigationList>
          <StyledNavigationList $align={ALIGN.center} />
          <StyledNavigationList $align={ALIGN.right}>
              <StyledNavigationItem>
                  {/* Stop default link behavior */}
                  <Button onClick={toggleDrawer} overrides={{ BaseButton: { style: { cursor: 'pointer' } } }}>TopPage</Button>
              </StyledNavigationItem>
          </StyledNavigationList>
      </HeaderNavigation>
    </MobileView>
    </>
  );
}
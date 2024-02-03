import React from "react";
import { Header } from "./Header.js";

export const NotFound = () => {
  return (
    <>
      <Header></Header>
      <header className='about-header'>
          <h1>404</h1>
      </header>
      <h3>お探しのページは見つかりませんでした。</h3>
      <p><a href="/">HOME</a></p>
    </>
  );
}
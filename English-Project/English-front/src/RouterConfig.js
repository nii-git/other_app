import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import { NotFound } from "./components/NotFound.js";
import { Provider } from "./components/Provider.js";
import  TopPage   from "./components/TopPage.js";
import  { About }  from "./components/About.js";
import {Providers} from "./components/Providers.js"


export const RouterConfig = ({isDarkMode,setDarkFunc}) => {
  return (
    <>
     <BrowserRouter>
      <Routes>
        <Route index element={<TopPage />} />
        {/* darkmodeいらないかも */}
        <Route path="provider" element={<Provider isDarkMode={isDarkMode} setDarkFunc={setDarkFunc} />} />
        <Route path="providers" element={<Providers isDarkMode={isDarkMode} setDarkFunc={setDarkFunc} />} />
        <Route path="about" element={<About />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
    </>
  );
}
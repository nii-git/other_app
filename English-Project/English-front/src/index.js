import React from 'react'
import ReactDOM from 'react-dom'
import './styles.css'
import { RouterConfig } from "./RouterConfig.js";
import { Client as Styletron } from "styletron-engine-monolithic";
import { Provider as StyletronProvider } from "styletron-react";
import { LightTheme, BaseProvider } from "baseui";




const engine = new Styletron();

ReactDOM.render(
  <React.StrictMode>
    <StyletronProvider value={engine}>
      <BaseProvider theme={LightTheme}>
        <RouterConfig />
      </BaseProvider>
    </StyletronProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
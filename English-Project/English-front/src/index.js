import React, { useEffect, useState } from "react";
import ReactDOM from 'react-dom'
import './styles.css'
import { RouterConfig } from "./RouterConfig.js";
import { Client as Styletron } from "styletron-engine-monolithic";
import { Provider as StyletronProvider } from "styletron-react";
import { LightTheme,DarkTheme, BaseProvider } from "baseui";

const Apps = () => {
  const [isDarkMode, setIsDarkMode] = useState(false);

  const engine = new Styletron();
  return (
    <React.StrictMode>
      <StyletronProvider value={engine}>
        {isDarkMode ? (
          <BaseProvider theme={DarkTheme}>
            <RouterConfig isDarkMode={isDarkMode} setDarkFunc={setIsDarkMode} />
          </BaseProvider>
        ) : (
          <BaseProvider theme={LightTheme}>
            <RouterConfig isDarkMode={isDarkMode} setDarkFunc={setIsDarkMode} />
          </BaseProvider>
        )}
      </StyletronProvider>
    </React.StrictMode>
  );
};

// Appsコンポーネントを直接レンダリングする
ReactDOM.render(<Apps />, document.getElementById('root'));

// const engine = new Styletron();


// ReactDOM.render(
//   <React.StrictMode>
//     <StyletronProvider value={engine}>
//       {isDarkMode? <BaseProvider theme={LightTheme}><RouterConfig />
//       </BaseProvider>: <BaseProvider theme={DarkTheme}><RouterConfig />
//       </BaseProvider>}
//     </StyletronProvider>
//   </React.StrictMode>,
//   document.getElementById('root')
// );
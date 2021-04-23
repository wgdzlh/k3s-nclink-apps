import { Admin, Resource, fetchUtils } from "react-admin";
import jsonServerProvider from "ra-data-json-server";
import AuthProvider from "./AuthProvider";

import {
  ModelList,
  ModelCreate,
  ModelEdit,
  ModelShow,
} from "./components/model";

import {
  AdapterList,
  AdapterCreate,
  AdapterEdit,
  AdapterShow,
} from "./components/adapter";

const httpClient = (url, options = {}) => {
  const auth = localStorage.getItem("auth");
  if (!auth) {
    return Promise.reject();
  }
  const { token } = JSON.parse(auth);
  options.user = {
    authenticated: true,
    token: `Bearer ${token}`,
  };
  return fetchUtils.fetchJson(url, options);
};

const dataProvider = jsonServerProvider(
  process.env.REACT_APP_API_URL,
  httpClient
);

function App() {
  return (
    <Admin authProvider={AuthProvider} dataProvider={dataProvider}>
      <Resource
        name="models"
        list={ModelList}
        show={ModelShow}
        create={ModelCreate}
        edit={ModelEdit}
      />
      <Resource
        name="adapters"
        list={AdapterList}
        show={AdapterShow}
        create={AdapterCreate}
        edit={AdapterEdit}
      />
    </Admin>
  );
}

export default App;

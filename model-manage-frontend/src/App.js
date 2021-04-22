import { Admin, Resource, fetchUtils } from "react-admin";
import jsonServerProvider from "ra-data-json-server";

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
  options.user = {
    authenticated: true,
    token:
      "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYWRtaW4ifQ.toNR4P8fvShGslzjTdOjQasLM-teflgYmeqEAz1NKJs",
  };
  return fetchUtils.fetchJson(url, options);
};

const dataProvider = jsonServerProvider(
  process.env.REACT_APP_API_URL,
  httpClient
);

function App() {
  return (
    <Admin dataProvider={dataProvider}>
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

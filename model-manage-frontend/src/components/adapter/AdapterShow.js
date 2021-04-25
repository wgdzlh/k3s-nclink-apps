import { Show, SimpleShowLayout, TextField, DateField } from "react-admin";

const AdapterShow = (props) => {
  return (
    <Show title="Show Adapter" {...props}>
      <SimpleShowLayout>
        <TextField source="id" />
        <DateField showTime source="created_at" />
        <DateField showTime source="updated_at" />
        <TextField source="dev_id" />
        <TextField source="model_id" />
      </SimpleShowLayout>
    </Show>
  );
};

export default AdapterShow;
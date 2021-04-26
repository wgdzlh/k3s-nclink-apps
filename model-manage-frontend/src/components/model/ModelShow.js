import {
  Show,
  SimpleShowLayout,
  TextField,
  DateField,
  NumberField,
  FunctionField,
  ReferenceManyField,
  SingleFieldList,
  ChipField,
} from "react-admin";

const ModelShow = (props) => {
  return (
    <Show title="Show Model" {...props}>
      <SimpleShowLayout>
        <TextField source="id" />
        <DateField showTime source="created_at" />
        <DateField showTime source="updated_at" />
        <NumberField source="used" />
        <TextField source="def" />
        <FunctionField
          label="Related adapters"
          render={(record) => {
            if (record.used === 0) return "None";
            return (
              <ReferenceManyField
                reference="adapters"
                target="model_id"
                sort={{ field: "id", order: "ASC" }}
              >
                <SingleFieldList linkType="show">
                  <ChipField source="id" />
                </SingleFieldList>
              </ReferenceManyField>
            );
          }}
        />
      </SimpleShowLayout>
    </Show>
  );
};

export default ModelShow;

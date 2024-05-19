import request from "@/utils/request";
import { categories, declarationTimes } from "@/utils/constants";

export async function postFilter(
  selectedCategories,
  selectedBatches,
  selectedEthnicity,
  selectedKeyword,
  selectedProvinces,
  updateHeatMap,
  updateCharts
) {
  let selectedCategoriesStrings = selectedCategories.value.map(
    (index) => categories[index]
  );
  if (selectedCategoriesStrings.length === 0) {
    selectedCategoriesStrings = ["all"];
  }

  let selectedBatchesStrings = selectedBatches.value.map(
    (index) => declarationTimes[index]
  );
  if (selectedBatchesStrings.length === 0) {
    selectedBatchesStrings = ["all"];
  }
  let selectedProvincesStrings = selectedProvinces.value;
  if (selectedProvincesStrings.length === 0) {
    selectedProvincesStrings = ["all"];
  }
  let selectedEthnicityStrings = selectedEthnicity.value.join(" ");
  const result = {
    category: selectedCategoriesStrings,
    batch: selectedBatchesStrings,
    ethnic: selectedEthnicityStrings,
    keyword: selectedKeyword.value,
    province: selectedProvincesStrings,
  };
  const jsonString = JSON.stringify(result);
  console.log(jsonString);
  await request.post("/filter", jsonString);
  updateHeatMap.value = true;
  updateCharts.value = true;
}

export function getHeat() {
  return request.get("/heat");
}

export function getChart(type) {
  return request.get("/chart", {
    params: {
      type: type,
    },
  });
}

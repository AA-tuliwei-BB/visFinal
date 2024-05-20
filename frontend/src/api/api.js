import request from "@/utils/request";
import { categories, declarationTimes } from "@/utils/constants";

export async function postFilter(
  selectedCategories,
  selectedBatches,
  selectedEthnicity,
  selectedKeyword,
  selectedProvinces,
  updateHeatMap,
  updateCharts,
  updateList
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
  updateList.value = true;
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

export async function getList(page, pageSize) {
  const response = await request.get("/list", {
    params: {
      page: page,
      size: pageSize,
    },
  });
  const declarationTimes = [
    "2006(第一批)",
    "2011(第三批)",
    "2021(第五批)",
    "2008(第二批)",
    "2014(第四批)",
  ];
  const myMap = new Map(
    declarationTimes.map((item) => {
      const batch = item.match(/\((.*?)\)/)[1];
      return [item, batch];
    })
  );

  response.data.forEach((item_1) => {
    if (myMap.has(item_1.batch)) {
      item_1.batch = myMap.get(item_1.batch);
    }
  });
  return response;
}

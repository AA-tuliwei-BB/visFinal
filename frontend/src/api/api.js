import request from "@/utils/request";

const categories = [
  "民间文学",
  "传统音乐",
  "传统舞蹈",
  "传统戏剧",
  "曲艺",
  "传统体育、游艺与杂技",
  "传统美术",
  "传统技艺",
  "传统医药",
  "民俗",
];

const declarationTimes = [
  "2006(第一批)",
  "2011(第三批)",
  "2021(第五批)",
  "2008(第二批)",
  "2014(第四批)",
];

export function postFilter(
  selectedCategories,
  selectedBatches,
  selectedEthnicity,
  selectedKeyword,
  selectedProvinces
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

  if (selectedProvinces.value.length === 0) {
    selectedProvinces.value = ["all"];
  }

  const result = {
    category: selectedCategoriesStrings,
    batch: selectedBatchesStrings,
    ethnic: selectedEthnicity.value,
    keyword: selectedKeyword.value,
    province: selectedProvinces.value,
  };

  const jsonString = JSON.stringify(result);
  console.log(jsonString);
  request.post("/filter", jsonString);
}

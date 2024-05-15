// store.js
import { createStore } from "vuex";

export default createStore({
  state: {
    selectedCategories: [],
    selectedEthnicity: [],
    selectedBatches: "",
    selectedKeyword: "",
    selectedProvinces: [],
  },
  mutations: {
    setSelectedCategories(state, category) {
      state.selectedCategories = category;
    },
    setSelectedEthnicity(state, ethnicity) {
      state.selectedEthnicity = ethnicity;
    },
    setSelectedBatches(state, batch) {
      state.selectedBatches = batch;
    },
    setSelectedKeyword(state, keyword) {
      state.selectedKeyword = keyword;
    },
    setSelectedProvinces(state, province) {
      state.selectedProvinces = province;
    },
  },
});

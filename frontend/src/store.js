// store.js
import { createStore } from "vuex";

export default createStore({
  state: {
    selectedCategories: [],
    selectedBatches: [],
    selectedEthnicity: "",
    selectedKeyword: "",
    selectedProvinces: [],
    updateHeatMap: false,
    updateCharts: false,
    updateList: false,
    updateRelationship: false,
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
    setUpdateHeatMap(state, update) {
      state.updateHeatMap = update;
    },
    setUpdateCharts(state, update) {
      state.updateCharts = update;
    },
    setUpdateList(state, update) {
      state.updateList = update;
    },
    setUpdateRelationship(state, update) {
      state.updateRelationship = update;
    },
  },
});

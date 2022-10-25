export const serializeNonPOJOs = (obj: Object) => {
  return JSON.parse(JSON.stringify(obj));
};

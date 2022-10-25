import { serializeNonPOJOs } from "../lib/utils";

export const load = ({ locals }: any) => {
  if (locals.user && locals.user.profile) {
    return {
      profile: serializeNonPOJOs(locals.user.profile)
    };
  };
}

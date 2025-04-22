import { ToastType } from "@/types/enums/toastTypeEnum";
import { toaster } from "./toaster";
// import { toaster } from "./toaster";



export function createToast(
    title?: string,
    description?: string,
    type?: ToastType,
    duration = 4000,
) {
    // const tPlacment: any = placement;

    toaster.create({
        title: title || "",
        description: description || "",
        duration,
        type: type || ToastType.Info,
    });
}

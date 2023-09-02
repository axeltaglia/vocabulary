export const getErrorMessage = (e: any): string => {
    if(typeof e.message === "string") {
        console.log(e)
        return e.message
    } else {
        throw e
    }
}
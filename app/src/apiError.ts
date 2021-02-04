export function showAPIError(errorMessage: string): void {
  (document.getElementById(
    "error-message"
  ) as HTMLElement).innerText = errorMessage;
  document.getElementById("error")?.classList.remove("hidden");
}

export function hideAPIError(): void {
  document.getElementById("error")?.classList.add("hidden");
}

export function showAPISuccess(successMessage: string): void {
  (document.getElementById(
    "success-message"
  ) as HTMLElement).innerText = successMessage;
  document.getElementById("success")?.classList.remove("hidden");
}

export function hideAPISuccess(): void {
  document.getElementById("success")?.classList.add("hidden");
}

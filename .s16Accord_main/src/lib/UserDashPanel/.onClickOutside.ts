/** Dispatch event on click outside of node */
export function clickOutside(node) {
  const handleClick = (event) => {
    if (node && !node.contains(event.target) && !event.defaultPrevented) {
      node.dispatchEvent(new CustomEvent("click_outside", node));
      window.removeEventListener('click', handleClick, true);
    }
  };  window.addEventListener('click', handleClick, true);
  return {
    destroy() {
      node.removeEventListener('click', handleClick);
    },
  };
}
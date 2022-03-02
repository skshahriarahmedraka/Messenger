import { writable } from "svelte/store";

export const count = writable(0)
export const showPeopleList = writable(0)
export const ChatOrDock = writable(0)
export const MessageList=writable([])
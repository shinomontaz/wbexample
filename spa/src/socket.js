import { io } from 'socket.io-client';
import {createContext, useState, useContext, React} from "react";

const ws_url = `${import.meta.env.VITE_WS_BASE_URL}`;

console.log("ws_url: " + ws_url);

export const socket = io(ws_url, {
  addTrailingSlash: false,
  path: '/ws',
  autoConnect: true,
  transports: ['websocket'],
});

export const SocketContext = createContext();

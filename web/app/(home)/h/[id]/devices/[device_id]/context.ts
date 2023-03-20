import { createContext } from 'use-context-selector'

interface IDeviceDetailContext {}

export const DeviceDetailContext = createContext<IDeviceDetailContext>({})

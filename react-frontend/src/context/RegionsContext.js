/* Region information is fetched once from the server and used in multiple places */
import React from 'react'

const RegionsContext = React.createContext()

export const RegionsProvider = RegionsContext.Provider
export const RegionsConsumer = RegionsContext.Consumer

export default RegionsContext
/* This Context is used for the Table Data. It is used for keeping TableSearch and Table components in sync */
import React from 'react'

const DataContext = React.createContext()

export const DataProvider = DataContext.Provider
export const DataConsumer = DataContext.Consumer

export default DataContext
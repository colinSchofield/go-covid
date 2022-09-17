import React from 'react'
import RegionsContext from '../../context/RegionsContext'

export default function Confirm({regions}) {
  const regionsContext = React.useContext(RegionsContext)

  const getFlag = (location) => {
    const country = location.trim()
    if (regionsContext.regionsData === null) {
      return ""
    }

    return regionsContext.regionsData
                  .filter((region) => region.key === country)
                  .map((element) => element.flag)
  }

  return (
    <>
    {regions.map((element) => (
        <span><span className="h4 emoji-align">{getFlag(element)}</span> {element}<br/></span>
      ))}
    </>
  )
}
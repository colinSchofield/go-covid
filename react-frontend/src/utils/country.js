import React from 'react'
import RegionsContext from '../context/RegionsContext'
import useGeoLocation from 'react-ipgeolocation'

export const useCountryFromGeoLocation = (setCountryCallback) => {
  const regionsContext = React.useContext(RegionsContext)
  const location = useGeoLocation()

  React.useEffect(() => {
    if (!location.isLoading && regionsContext.regionsData) {
      const region = regionsContext.regionsData
                      .find((region) => region.countryCode === location.country)
      if (region) {
        setCountryCallback(region.key)
      }
    }
  }, [location])
}

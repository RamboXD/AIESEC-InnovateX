import React from 'react'

interface HomeLayoutProps {
  children: React.ReactNode
}

const HomeLayout: React.FC<HomeLayoutProps> = ({children}) => {
  return (
    <div className='w-full h-full px-4 py-4 flex flex-col gap-3'>{children}</div>
  )
}

export default HomeLayout
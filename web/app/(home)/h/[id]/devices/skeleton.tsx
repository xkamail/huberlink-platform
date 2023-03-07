const DeviceSkeletons = () => {
  return (
    <>
      {[...new Array(3)].map((_, i) => (
        <div
          className="col-span-1 flex justify-between items-center rounded-lg p-4 bg-white/50 shadow transition-all cursor-pointer hover:shadow-lg h-[72px]"
          key={i}
        >
          <div></div>
        </div>
      ))}
    </>
  )
}

export default DeviceSkeletons

DeviceSkeletons.displayName = 'DeviceSkeletons'

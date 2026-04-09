import Navbar from '../../components/Navbar';
import useAPISetting from '../../hooks/APISetting';

export default function Konfigurasi() {
  const { baseUrl, setBaseUrl, handlePostApiSetting, apiData } = useAPISetting();

  return (
    <div className="min-h-screen bg-gray-50 font-poppins-regular py-24">
      <Navbar/>

      <div className="max-w-7xl mx-auto px-6 py-8" data-aos="fade-up" data-aos-duration="1000">
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 mb-6">
          <h1 className="text-2xl font-poppins-bold text-gray-800">
            Konfigurasi Website
          </h1>
        </div>

        <div className="bg-white rounded-xl shadow-lg mb-8 overflow-hidden border border-gray-100">
          <div className="bg-linear-to-r from-third to-secondary px-6 py-5 relative overflow-hidden">
            <div className="absolute inset-0 bg-linear-to-br from-white/10 to-transparent"></div>
            <div className="absolute -right-10 -top-10 w-40 h-40 bg-white/10 rounded-full blur-3xl"></div>
            <div className="absolute -left-10 -bottom-10 w-40 h-40 bg-white/10 rounded-full blur-3xl"></div>
            <div className="relative flex items-center justify-between">
              <div className="flex items-center gap-3">
                <div className="w-10 h-10 bg-white/20 backdrop-blur-sm rounded-xl flex items-center justify-center shadow-lg">
                  <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                </div>
                <div>
                  <h2 className="text-xl font-poppins-bold text-white">Konfigurasi API</h2>
                  <p className="text-hover text-xs font-poppins-regular mt-0.5">Atur koneksi API</p>
                </div>
              </div>
              {apiData.length  && (
                <div className="flex items-center gap-2 bg-white/20 backdrop-blur-sm px-3 py-1.5 rounded-lg">
                  <svg className="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <span className="text-xs font-poppins-medium text-white">Tersimpan</span>
                </div>
              )}
            </div>
          </div>

          <div className="p-8" data-aos="fade-up" data-aos-duration="1000">
            <div className="grid grid-cols-1 gap-6 mb-6">
              <div className="group">
                <label className="flex items-center gap-2 text-sm font-poppins-semibold text-gray-700 mb-3">
                  <svg className="w-4 h-4 text-secondary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9" />
                  </svg>
                  Base URL
                  <span className="text-red-500">*</span>
                </label>
                <div className="relative">
                  <div className="absolute left-4 top-1/2 -translate-y-1/2 pointer-events-none">
                    <svg className="w-5 h-5 text-gray-400 group-focus-within:text-third transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
                    </svg>
                  </div>
                  <input
                    type="url"
                    value={baseUrl}
                    onChange={(e) => setBaseUrl(e.target.value)}
                    className="w-full pl-12 pr-4 py-3.5 border-2 border-gray-200 rounded-xl font-poppins-regular
                      focus:border-third focus:ring-4 focus:ring-hover outline-none
                      transition-all duration-300 hover:border-gray-300
                      bg-linear-to-br from-white to-gray-50"
                    placeholder="api.monalisa.example.com"
                  />
                </div>
              </div>
            </div>

            <div className="flex items-center justify-between pt-4 border-t border-gray-100">
              <div className="flex items-center gap-2 text-sm text-gray-600">
                <svg className="w-4 h-4 text-third" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span className="font-poppins-regular">Konfigurasi ini akan terhubung ke sistem</span>
              </div>
              <div className="flex gap-3">
                <button
                  type="submit"
                  onClick={() => handlePostApiSetting()}
                  disabled={!baseUrl}
                  className="flex items-center gap-2 px-6 py-2.5 bg-linear-to-r from-third to-secondary hover:from-secondary hover:to-secondary text-white rounded-lg font-poppins-semibold transition-all duration-200 shadow-lg shadow-hover hover:shadow-xl disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:shadow-hover"
                >
                  <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4" />
                  </svg>
                  {apiData[0] ? 'Update Konfigurasi' : 'Simpan Konfigurasi'}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
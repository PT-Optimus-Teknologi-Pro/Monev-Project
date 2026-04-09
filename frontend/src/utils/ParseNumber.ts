export const ParseNumber = (value: any): number => {
  if (value == null || value === '') return 0;

  if (typeof value === 'number') {
    if (value > 0 && value < 1) {
      return value * 100;
    }
    return value;
  }

  const str = String(value).trim();
  if (str.includes('%')) {
    return Number(str.replace(/[^0-9.-]/g, '')) || 0;
  }

  return Number(str.replace(/[^0-9.-]/g, '')) || 0;
};
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

export const ParseIDRupiahNumber = (value: any): number => {
  if (value == null || value === '') return 0;

  if (typeof value === 'number') {
    return value;
  }

  let str = String(value).trim();
  str = str.replace(/\s+/g, '');

  str = str.replace(/[^0-9,.-]/g, '');

  if (!str) return 0;

  if (str.includes(',')) {
    str = str.replace(/\./g, '').replace(',', '.');
  } else {
    str = str.replace(/\./g, '');
  }

  const num = Number(str);
  if (isNaN(num)) return 0;
  return num;
};
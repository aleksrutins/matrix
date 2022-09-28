export type Config = {
  configPath: string
  move: {
    from: string
    to: string
  }[]
  configRegex: {
    find: string
    replace: string
  }
  configurations: Record<string, string>
  targets: Record<string, string[]>
}

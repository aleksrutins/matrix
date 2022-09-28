import * as core from '@actions/core'
import TOML from '@ltd/j-toml'
import * as fs from 'fs/promises'
import {Config} from './config'
import * as childProcess from 'child_process'

async function run(): Promise<void> {
  const config = TOML.parse(
    await fs.readFile(core.getInput('config'))
  ) as Config
  const currentConfig = (await fs.readFile(config.configPath)).toString()
  const regex = new RegExp(config.configRegex.find, 'g')
  for (const cfg in config.configurations) {
    const replacement = config.configRegex.replace.replace(
      /\$\(config\)/g,
      config.configurations[cfg]
    )
    const newConfig = currentConfig.replace(regex, replacement)
    await fs.writeFile(config.configPath, newConfig)

    core.startGroup(cfg)
    for (const target in config.targets) {
      core.startGroup(`Target ${target}`)
      childProcess.spawnSync(
        config.targets[target][0],
        config.targets[target].slice(1),
        {stdio: 'inherit'}
      )
      core.endGroup()
    }
    core.endGroup()

    for (const file of config.move) {
      await fs.rename(file.from, file.to)
    }
  }
}

run()
